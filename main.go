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
	Content     string
	Date        time.Time
	Category    string
	Excerpt     string
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
	s.templates.ExecuteTemplate(w, "home.html", data)
}

func (s *Server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "About - Charlotte's Portfolio",
	}
	s.templates.ExecuteTemplate(w, "about.html", data)
}

func (s *Server) projectsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Projects - Charlotte's Portfolio",
	}
	s.templates.ExecuteTemplate(w, "projects.html", data)
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
			
			// Parse markdown content
			htmlContent := parser.MarkdownToHTML(content)
			
			// Extract metadata from content
			lines := strings.Split(string(content), "\n")
			var title, category, dateStr string
			var contentStart int
			
			for i, line := range lines {
				if strings.HasPrefix(line, "# ") {
					title = strings.TrimPrefix(line, "# ")
					contentStart = i + 1
				} else if strings.Contains(line, "*Published:") {
					dateStr = strings.TrimSpace(strings.Split(line, ":")[1])
					dateStr = strings.Trim(dateStr, "*")
				} else if strings.Contains(line, "*Category:") {
					category = strings.TrimSpace(strings.Split(line, ":")[1])
					category = strings.Trim(category, "*")
				}
			}
			
			// Parse date
			date, err := time.Parse("January 2, 2006", strings.TrimSpace(dateStr))
			if err != nil {
				date = time.Now()
			}
			
			// Create slug from filename
			slug := strings.TrimSuffix(filepath.Base(path), ".md")
			
			// Extract excerpt (first paragraph)
			excerpt := ""
			if contentStart < len(lines) {
				for i := contentStart; i < len(lines); i++ {
					if strings.TrimSpace(lines[i]) != "" && !strings.HasPrefix(lines[i], "*") {
						excerpt = strings.TrimSpace(lines[i])
						break
					}
				}
			}
			
			post := BlogPost{
				Title:    title,
				Slug:     slug,
				Content:  string(htmlContent),
				Date:     date,
				Category: category,
				Excerpt:  excerpt,
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
	http.HandleFunc("/about", server.aboutHandler)
	http.HandleFunc("/projects", server.projectsHandler)
	http.HandleFunc("/contact", server.contactHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
