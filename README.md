# Go-Course

This repository is my personal journey of learning Go, featuring a curated set of practical projects and exercises. Each folder documents a key concept, pattern, or real-world use case I explored while mastering Go and its integration with web technologies.

## What I Learned

- **Femm (Frontend Masters Museum)**  
  My first web app in Go! I built a museum showcase with both backend and frontend:
  - Created REST API endpoints for listing and adding exhibitions (`/api/exhibitions`, `/api/exhibitions/add`).
  - Designed a dynamic frontend using HTML, CSS, and JavaScript to render exhibitions like “Life in Ancient Greek”, “Aristotle: Life and Legacy”, and more.
  - Experimented with serving static files, templates, and using Gin for routing.

- **FEM**  
  Here, I learned Go's approach to object-oriented programming:
  - Defined structs for instructors, courses, and workshops.
  - Implemented interfaces and practiced struct embedding.
  - Explored how interfaces can simplify logic (e.g., sign-up methods).

- **GoRoutines**  
  My introduction to concurrency:
  - Created basic goroutines and channels.
  - Observed how Go handles parallel tasks and communication.

- **Files**  
  Learned practical file I/O:
  - Read from and write to files.
  - Handled errors and manipulated file contents.

- **Input_Output**  
  Practiced Go fundamentals:
  - Used constants, pointers, and basic I/O.
  - Built and debugged simple functions.

- **Calculator**  
  My first console-based Go project:
  - Implemented a simple calculator with arithmetic operations via CLI.

- **CryptoMasters**  
  Combined concurrency with HTTP APIs:
  - Fetched cryptocurrency rates concurrently.
  - Processed and displayed results in parallel.

## Directory Structure

- `Femm/` - Museum web app (Go backend, HTML/JS/CSS frontend)
- `FEM/` - Go OOP examples (courses & workshops)
- `GoRoutines/` - Concurrency demos
- `files/` - File I/O utilities
- `Input_Output/` - Basic Go syntax/IO
- `calculator/` - CLI calculator
- `cryptomasters/` - Concurrent crypto price fetcher

## Technologies Used

- Go (main language)
- HTML, CSS, JavaScript (frontend for museum app)
- Gin (for some web routing)
- Native Go net/http

---

If you're looking to follow a similar path or just want practical Go examples, feel free to browse the code!
