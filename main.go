package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Welcome Hendrich</title>
				<style>
					@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');
					
					:root {
						--primary-color: #8e44ad;
						--secondary-color: #e74c3c;
						--accent-color: #f1c40f;
						--text-color: #ffffff;
						--bg-gradient: linear-gradient(135deg, #9b59b6 0%, #6a0dad 100%);
					}
					
					* {
						box-sizing: border-box;
						margin: 0;
						padding: 0;
					}
					
					body {
						background: var(--bg-gradient);
						font-family: 'Poppins', sans-serif;
						margin: 0;
						padding: 0;
						display: flex;
						justify-content: center;
						align-items: center;
						min-height: 100vh;
						color: var(--text-color);
						overflow-x: hidden;
					}
					
					.container {
						background-color: rgba(255, 255, 255, 0.15);
						backdrop-filter: blur(10px);
						border-radius: 20px;
						padding: 50px;
						box-shadow: 0 15px 35px rgba(0, 0, 0, 0.3);
						text-align: center;
						max-width: 800px;
						width: 90%;
						animation: fadeIn 1.2s ease-in, float 6s ease-in-out infinite;
						border: 1px solid rgba(255, 255, 255, 0.2);
					}
					
					h1 {
						color: var(--text-color);
						font-size: 4rem;
						margin-bottom: 25px;
						text-shadow: 3px 3px 6px rgba(0, 0, 0, 0.3);
						letter-spacing: 1px;
					}
					
					p {
						color: var(--text-color);
						font-size: 1.5rem;
						line-height: 1.7;
						margin-bottom: 30px;
						opacity: 0.9;
					}
					
					.highlight {
						color: var(--accent-color);
						font-weight: 700;
						text-shadow: 1px 1px 3px rgba(0, 0, 0, 0.3);
						position: relative;
						display: inline-block;
					}
					
					.highlight::after {
						content: '';
						position: absolute;
						bottom: -5px;
						left: 0;
						width: 100%;
						height: 3px;
						background-color: var(--accent-color);
						border-radius: 10px;
					}
					
					.button-container {
						display: flex;
						justify-content: center;
						gap: 20px;
						margin-top: 40px;
					}
					
					.button {
						background-color: var(--secondary-color);
						color: white;
						border: none;
						padding: 15px 30px;
						font-size: 1.2rem;
						border-radius: 50px;
						cursor: pointer;
						transition: all 0.4s ease;
						text-decoration: none;
						display: inline-block;
						font-weight: 600;
						box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
					}
					
					.button:hover {
						transform: translateY(-5px) scale(1.05);
						box-shadow: 0 10px 20px rgba(0, 0, 0, 0.3);
					}
					
					.button.primary {
						background-color: var(--secondary-color);
					}
					
					.button.secondary {
						background-color: transparent;
						border: 2px solid var(--text-color);
					}
					
					.tag {
						display: inline-block;
						background-color: rgba(255, 255, 255, 0.2);
						padding: 8px 16px;
						border-radius: 50px;
						font-size: 0.9rem;
						margin-bottom: 30px;
						backdrop-filter: blur(5px);
						border: 1px solid rgba(255, 255, 255, 0.1);
					}
					
					@keyframes fadeIn {
						from { opacity: 0; transform: translateY(-30px); }
						to { opacity: 1; transform: translateY(0); }
					}
					
					@keyframes float {
						0% { transform: translateY(0px); }
						50% { transform: translateY(-15px); }
						100% { transform: translateY(0px); }
					}
					
					.divider {
						width: 80%;
						height: 1px;
						background: linear-gradient(to right, transparent, rgba(255, 255, 255, 0.5), transparent);
						margin: 30px auto;
					}
					
					.features {
						display: flex;
						justify-content: space-around;
						margin: 30px 0;
						flex-wrap: wrap;
					}
					
					.feature {
						flex: 1;
						min-width: 200px;
						padding: 15px;
						margin: 10px;
					}
					
					.feature-icon {
						font-size: 2rem;
						margin-bottom: 15px;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<span class="tag">K8S with GitOps</span>
					<h1>Hello, <span class="highlight">Hendrich</span>!</h1>
					<p>Welcome to this personalized demo application created just for you.</p>
					<p>This is a sample project showcasing a modern web interface with Go backend. Perfect for GitOps demonstrations and containerized applications.</p>
					
					<div class="divider"></div>
					
					<div class="features">
						<div class="feature">
							<div class="feature-icon">🚀</div>
							<p>Fast & Lightweight</p>
						</div>
						<div class="feature">
							<div class="feature-icon">🔒</div>
							<p>Secure Design</p>
						</div>
						<div class="feature">
							<div class="feature-icon">⚙️</div>
							<p>Easy to Deploy</p>
						</div>
					</div>
					
					<div class="button-container">
						<a href="#" class="button primary">Get Started</a>
						<a href="#" class="button secondary">Learn More</a>
					</div>
				</div>
			</body>
			</html>
		`))
	})

	port := ":1337"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to stop the server")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down server...")
}
