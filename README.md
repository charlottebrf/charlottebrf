# Charlotte's Portfolio Website

A personal portfolio website built with Go, featuring a blog with markdown support and a dark ocean-inspired design.
This site replaces a [previous version written using 11ty](https://github.com/charlottebrf/charlottebrf.dev).
## 🌐 Live Site

**Production:** [charlotte-go-website.fly.dev](https://charlotte-go-website.fly.dev)

## ✨ Features

- **Dark Theme**: Ocean-inspired design with turquoise accents
- **Blog System**: Markdown-powered blog with status tags (draft/complete)
- **Experience Timeline**: Professional background showcase
- **Responsive Design**: Mobile-friendly layout
- **Analytics**: Privacy-focused tracking with Plausible
- **Fast Performance**: Lightweight Go server

## 🛠 Tech Stack

- **Backend**: Go with html/template
- **Markdown**: Custom parser with github.com/gomarkdown/markdown
- **Styling**: CSS with Grid and Flexbox
- **Fonts**: Inconsolata (monospace), Lora (serif)
- **Deployment**: Fly.io
- **Analytics**: Plausible

## 📁 Project Structure

```
charlotte-go-website/
├── content/
│   └── blog/           # Markdown blog posts
├── static/
│   ├── css/           # Stylesheets
│   ├── images/        # Static images
│   └── js/            # JavaScript files
├── templates/         # HTML templates
├── parser/            # Markdown parsing logic
├── main.go           # Web server
├── go.mod            # Go dependencies
└── Dockerfile        # Container configuration
```

## 🚀 Local Development

### Prerequisites

- Go 1.24+
- Git

### Setup

1. **Clone the repository**
   ```bash
   git clone git@github.com:charlottebrf/charlotte-go-website.git
   cd charlotte-go-website
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

4. **Visit the site**
   Open [http://localhost:8080](http://localhost:8080)

## 📝 Adding Blog Posts

Create a new markdown file in `content/blog/` with the following format:

```markdown
# Your Post Title

*Published: July 1, 2025*
*Category: Technology*
*Status: draft*

Your blog post content here...
```

**Supported fields:**
- `Published`: Date in "Month Day, Year" format
- `Category`: Post category
- `Status`: `draft` or `complete` (shows colored tag)

## 🎨 Design System

### Colors
- **Background**: `#0d1b2a` (dark ocean blue)
- **Primary**: `#20b2aa` (turquoise)
- **Text**: `#eceff1` (light gray)
- **Secondary**: `#707070` (medium gray)

### Typography
- **Headers**: Lora (serif)
- **Body**: Lora (serif)
- **Code/Nav**: Inconsolata (monospace)

## 🚢 Deployment

The site is automatically deployed to Fly.io. To deploy manually:

```bash
fly deploy
```
There is a custom domain setup.

## 📊 Analytics

Analytics are provided by Plausible (privacy-focused, GDPR compliant). The tracking script is included in all page templates with the domain `charlotte-go-website.fly.dev`.

## 🔧 Environment Variables

- `PORT`: Server port (defaults to 8080)

## 📄 License

© 2025 Charlotte Fereday. All rights reserved.
See the [MIT License](LICENSE) relevant for the code and [CC by NC ND License for the blog content](LICENSE-CC-BY-NC-ND).