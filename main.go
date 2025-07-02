package main

import (
	"charlotte-go-website/parser"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Server struct {
	templates *template.Template
}

type BlogPost struct {
	Title       string
	Slug        string
	Content     template.HTML
	Date        time.Time
	Category    string
	Excerpt     string
	Status      string
}

type BlogData struct {
	Title string
	Posts []BlogPost
}

func (s *Server) loadTemplates() {
	s.templates = template.Must(template.ParseGlob("templates/*.html"))
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Charlotte's Portfolio",
	}
	s.templates.ExecuteTemplate(w, "about.html", data)
}


func (s *Server) loadBlogPosts() ([]BlogPost, error) {
	var posts []BlogPost
	blogDir := "content/blog"
	
	err := filepath.WalkDir(blogDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			
			// Extract metadata from content
			lines := strings.Split(string(content), "\n")
			var title, category, dateStr, status string
			var contentStart int
			
			for i, line := range lines {
				if strings.HasPrefix(line, "# ") {
					title = strings.TrimPrefix(line, "# ")
				} else if strings.Contains(line, "*Published:") {
					dateStr = strings.TrimSpace(strings.Split(line, ":")[1])
					dateStr = strings.Trim(dateStr, "*")
				} else if strings.Contains(line, "*Category:") {
					category = strings.TrimSpace(strings.Split(line, ":")[1])
					category = strings.Trim(category, "*")
				} else if strings.Contains(line, "*Status:") {
					status = strings.TrimSpace(strings.Split(line, ":")[1])
					status = strings.Trim(status, "*")
				} else if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "*") && !strings.HasPrefix(line, "# ") && contentStart == 0 {
					// First non-metadata line marks the start of content
					contentStart = i
				}
			}
			
			// Default status to "complete" if not specified
			if status == "" {
				status = "complete"
			}
			
			// Create content without metadata for markdown parsing
			var contentLines []string
			if contentStart > 0 {
				contentLines = lines[contentStart:]
			} else {
				contentLines = lines
			}
			markdownContent := strings.Join(contentLines, "\n")
			
			// Parse markdown content
			htmlContent := parser.MarkdownToHTML([]byte(markdownContent))
			
			// Parse date
			date, err := time.Parse("January 2, 2006", strings.TrimSpace(dateStr))
			if err != nil {
				date = time.Now()
			}
			
			// Create slug from filename
			slug := strings.TrimSuffix(filepath.Base(path), ".md")
			
			// Extract excerpt (first paragraph)
			excerpt := ""
			if contentStart > 0 && contentStart < len(lines) {
				excerpt = strings.TrimSpace(lines[contentStart])
			}
			
			post := BlogPost{
				Title:    title,
				Slug:     slug,
				Content:  template.HTML(htmlContent),
				Date:     date,
				Category: category,
				Excerpt:  excerpt,
				Status:   status,
			}
			
			posts = append(posts, post)
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Sort posts by date (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	
	return posts, nil
}

func (s *Server) blogHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.loadBlogPosts()
	if err != nil {
		http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
		return
	}
	
	data := BlogData{
		Title: "Blog - Charlotte's Portfolio",
		Posts: posts,
	}
	
	s.templates.ExecuteTemplate(w, "blog.html", data)
}

func (s *Server) blogPostHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	if slug == "" || slug == "blog" {
		http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
		return
	}
	
	posts, err := s.loadBlogPosts()
	if err != nil {
		http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
		return
	}
	
	var post *BlogPost
	for _, p := range posts {
		if p.Slug == slug {
			post = &p
			break
		}
	}
	
	if post == nil {
		http.NotFound(w, r)
		return
	}
	
	data := struct {
		Title string
		Post  *BlogPost
	}{
		Title: post.Title + " - Charlotte's Portfolio",
		Post:  post,
	}
	
	s.templates.ExecuteTemplate(w, "blog-post.html", data)
}

func (s *Server) experienceHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Experience - Charlotte's Portfolio",
	}
	s.templates.ExecuteTemplate(w, "experience.html", data)
}

func (s *Server) contactHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Contact - Charlotte's Portfolio",
	}
	s.templates.ExecuteTemplate(w, "contact.html", data)
}

func main() {
	server := &Server{}
	server.loadTemplates()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/blog/", server.blogPostHandler)
	http.HandleFunc("/blog", server.blogHandler)
	http.HandleFunc("/", server.homeHandler)
	http.HandleFunc("/experience", server.experienceHandler)
	http.HandleFunc("/contact", server.contactHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
