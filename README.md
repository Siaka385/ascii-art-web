# ASCII Art Web

## Description

A web application that converts input text to ASCII art using different banner styles.The project also contain dockerfile for dockerization

Docker is a platform that allows you to run applications in isolated environments called containers. A Docker container is a lightweight, self-contained package that includes everything needed to run an application, such as the code, libraries, and settings. Containers are created from Docker images, which are templates containing all the components required for a project. These images are built using a Dockerfile, a script that outlines the steps to set up the application, like installing dependencies and copying files.
</br>
  <section class="authors">
                <h1> AUTHORS :</h1> </br>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/fb9713e670165fd5fc7536ccc10a6c8d?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/tesiaka">tesiaka</a></p>
                </div>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/27e7d22c36267615242d9bc4a275d831?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/jerootieno">jerootieno</a></p>
                </div>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/cf0006d1b23256772956a4629c7a25a1?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/viarony">viarony</a></p>
                </div>
            </section>

## Usage

To run the application:

1. Clone the repository:
    ```sh
    git clone https://github.com/Siaka385/ascii-art-web.git
    ```
2. Navigate to the project directory:
    ```sh
    cd ascii-art-web
    ```
3. Build and run the application:
    ```sh
    go run .
    ```
4. Open your web browser and navigate to `http://localhost:8080`.

## Implementation Details

### Algorithm

1. **Input Handling**:
    - The input text and banner style are received via an HTML form.
    - The text is split by newlines if it contains them.

2. **ASCII Art Generation**:
    - The application reads the specified banner file.
    - It converts each character of the input text to its ASCII art representation.
    - It validates that the input contains only printable ASCII characters.

3. **Error Handling**:
    - Returns appropriate HTTP status codes:
        - `200 OK`: If everything went without errors.
        - `404 Not Found`: If templates or banners are not found.
        - `400 Bad Request`: For incorrect requests.
        - `500 Internal Server Error`: For unhandled errors.
        - `405 method not allowed`:For wrong method used

## Docker Functionality

### Installing Docker

To install Docker on your system, follow the instructions based on your operating system:

#### For Ubuntu/Debian:

1.Update your package index:


```sh
sudo apt-get update
```


2.Install required packages:

```sh
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
```
3.Add Docker’s official GPG key:

```sh
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

4.Add Docker’s APT repository:

```sh
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

5.Update the package index again:


```sh
sudo apt-get update
```

6.Install Docker CE (Community Edition):

```sh
sudo apt-get install docker-ce
```

7. Verify that Docker is installed and running:

```sh
    sudo systemctl status docker
```

#### For macOS:

-  Download Docker Desktop for Mac from the Docker website.

- Open the downloaded .dmg file and drag the Docker icon to the Applications folder.

- Open Docker from the Applications folder and follow the installation instructions.

    Verify Docker is installed by running:

```sh
    docker --version
```

#### For Windows:

- Download Docker Desktop for Windows from the Docker website.

- Run the installer and follow the installation instructions.

- After installation, launch Docker Desktop from the Start menu.

- Verify Docker is installed by running:

```sh
    docker --version
```

### Building a Docker Image in this project

To build a Docker image for the application, use the following command from the root of the project directory where the Dockerfile is located:

```sh
docker build -t ascii-art-web .
```

This command tags the image as ascii-art-web.

### Running a Docker Container

To run a Docker container from the image:

```sh

docker run -p 8086:8080 --name ascii-art-container ascii-art-web
```
- -p 8086:8080 maps port 8080 in the container to port 8086 on your host.
- --name ascii-art-container assigns a name to the container.

### Removing a Docker Image

To remove a Docker image, use:

```sh

docker rmi <image_id>
```
Replace <image_id> with the actual ID of the image you want to remove. To list all images and find the image ID, use:

```sh

docker images
```

### Checking the Number of Docker Images

To check the number of Docker images on your system:

```sh

docker images -q | wc -l
```
This command lists all image IDs and counts them.

### Creating a Docker Container

To create and start a Docker container, use the docker run command as shown above in "Running a Docker Container."

### Stopping a Docker Container

To stop a running container:

```sh

docker stop <container_id>
```
Replace <container_id> with the ID or name of the container. You can find the container ID using:

```sh

docker ps
```

### Removing a Docker Container

To remove a stopped container:

```sh

docker rm <container_id>
```
Replace <container_id> with the ID or name of the container.

### Entering Inside a Docker Container

To open an interactive shell inside a running container:

```sh

docker exec -it <container_id> sh
```
Replace <container_id> with the ID or name of the container. If your container uses a different shell, such as bash, adjust the command accordingly:

```sh

docker exec -it <container_id> bash
```

### Listing Files Inside a Docker Container

Once inside the container, you can list files and directories with detailed information using:

```sh

ls -l
```
This command shows a detailed list of files and directories in the current working directory, including permissions, ownership, size, and modification date.


## Contributors
feel free to fork this repo and contribute
