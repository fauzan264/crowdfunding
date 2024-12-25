# BWA Crowdfunding

## Description
This project is based on the **Build with Angga** learning flow, enhanced with custom improvisations. It involves developing a full-stack web application with a Vue.js-based front-end, a Golang-based back-end, and a MySQL database.

### **Learning Flow**
1. **Slicing HTML to Vue Template**  
   - Create the user interface using HTML and CSS, then integrate it into Vue.js components for a dynamic and responsive layout.

2. **Back-End Development with Golang**  
   - Build efficient and scalable APIs using Golang to connect the front-end and the database.

3. **Front-End Development with NuxtJS**  
   - Develop the front-end using NuxtJS, a Vue-based framework, for a modern, server-side rendered (SSR) experience.

4. **Personal Improvisations**  
   - Added custom features and optimizations to improve functionality, user experience, and performance.

---

## Technologies Used
- **HTML/CSS**: For slicing and styling the front-end layout.  
- **Vue.js**: Version `3.5.13` for creating modular UI components.  
- **Golang**: Version `1.22.4` for building the back-end API.  
- **NuxtJS**: Version `3.14.1592` for creating a modern and SEO-friendly front-end application.  
- **MySQL**: Version `8.0` as the database solution.  

---

## Project Requirements
### **Software Requirements**
1. **Docker & Docker Compose**  
   - Docker is used to containerize the services for easier development and deployment.

2. **Node.js** (v18.x or later)  
   - Required for running the NuxtJS frontend container.

3. **Go** (v1.22.4 or later)  
   - Required for running the Golang backend container.

4. **Database**  
   - The project uses **MySQL 8.0**, managed through Docker.

---

## How to Run the Project with Docker

### 1. Clone the Repository
```bash
git clone https://github.com/fauzan264/crowdfunding.git
cd crowdfunding
```

### 2. Create Environment Files
- Create a `.env` file in the **root directory** by copying the example file:
```bash
cp .env-example .env
```
- Do the same for the backend directory:
```bash
cp  backend/.env-example backend/.env
```

### 3. Start Docker
Ensure Docker is installed and running, then execute the following command:
```bash
docker-compose up --build
```

### 4. Wait for the Setup to Complete
Docker will handle the following:
- Downloading the necessary images
- Building the containers
- Starting the services
This process might take a few minutes. Once completed, the application will be ready to use.

### 5. Access the Application
- Frontend: Open your browser and navigate to http://localhost:3000.
- Backend API: Access the API at http://localhost:8080.