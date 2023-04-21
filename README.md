# MarketMe Web Server
## Containerized golang web application


* The included Dockerfile is building the application using multi-stage building.
* The end result image in no more than 15mb thanks to the "Scratch" docker image which is the smallest image out there.
* a Jenkinsfile is also included in order to build the application and push it to my private DockerHub.  
* My home lab Jenkins Server runs a pipeline to build the golang application when a push is detected using the Poll SCM option.

  <img src="webapp/images/pipeline.jpg" alt="alt text" width="400" height="60">

### __How to run the application in localhost__:
1. clone the repository
```
git clone -b go-webapp https://github.com/tomerschwartz24/Go-Projects.git
```
2. cd to Go-Projects directory
```
cd Go-Projects
```
3. Build the application image 
```
docker build -t marketme:1.0 -f Dockerfile .
```
4. Run the built image 
```
docker run -p 80:80 -d marketme:1.0 
```