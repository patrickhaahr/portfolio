# Portfolio Website
![PORTFOLIO](https://github.com/user-attachments/assets/bd372c5c-6781-430b-882e-64b0d29d7642)

## Overview

This is my personal portfolio website designed to showcase my skills, experience, and projects. It features a responsive design with a dark mode toggle, an interactive tech stack section, and a contact form that sends emails directly to me. The website is built using Go, HTMX, Templ, and TailwindCSS for clean, efficient, and scalable performance.

## Features

-   **Responsive Design:** Works across all devices, including mobile and desktop.
-   **Dark Mode:** Users can toggle between light and dark themes.
-   **Contact Form:** Allows users to send messages directly to my email.
-   **Tech Stack:** An interactive display of technologies I use, with swiping functionality on smaller screens.
-   **Project Showcase:** Highlights several projects I've worked on, with direct links to the demo and source code.

## Technologies Used

-   **Backend:** Go, HTMX, Templ
-   **Frontend:** TailwindCSS, HTMX
-   **Email Integration:** SendGrid API for the contact form
-   **Deployment:** Heroku

## Getting Started

### Prerequisites

-   [Go](https://golang.org/doc/install)
-   Node.js (for Tailwind CSS)
-   SendGrid API Key

### Setup Instructions

Clone the project and install all its dependancies:

    
    git clone https://github.com/patrickhaahr/portfolio.git
    cd portfolio
    go mod tidy
    npm install
    

Set up your environment variables for SendGrid:

    export SENDGRID_API_KEY=your-sendgrid-api-key

To run the project locally, use the provided Makefile commands:

  -   To watch and compile CSS with Tailwind:

    make css

  -   To generate templ files and enable live reloading using templ:

    make templ

Use Air for live reloading your Go server. Make sure you have the ".air.toml" configuration set up, then simply run:

    air

This will start the development server with live reloading, compiling your templates and CSS in real-time as you make changes.

### Customization

-   **Update Resume Documents**: Replace the existing PDF and update the resume link in the navigation.
-   **Update External Links**: Modify GitHub, LinkedIn, and other links in the navigation and footer components.
-   **Change Contact Email**: Update the email handler and API keys for form submissions.
-   **Modify Tech Stack Icons**: Add, remove, or replace SVG icons in the tech stack.
-   **Change Project Links**: Edit project titles, descriptions, and links to match your own projects and demos.

## License

This project is licensed under the MIT License.

* * *

Feel free to use this project as inspiration for your own portfolio! If you have any questions or feedback, don't hesitate to contact me via the contact form on the website.
