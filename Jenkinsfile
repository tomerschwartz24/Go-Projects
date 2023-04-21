pipeline {
    agent any
    stages {
        stage('Building Website') {
            steps {
                sh 'docker build -t tomerschwartz2411/website:webapp_$BUILD_NUMBER -f Dockerfile .'
            }
        }
        
        stage('Pushing to DockerHub') {
            steps {
                withCredentials([usernamePassword( credentialsId: 'docker_hub', usernameVariable: 'USER', passwordVariable: 'PASSWORD')]) {
                sh 'docker login -u $USER -p $PASSWORD docker.io'
                sh 'docker push tomerschwartz2411/website:webapp_$BUILD_NUMBER'
                } 
            }
        }
    }
}   
