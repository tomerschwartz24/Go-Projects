pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/go-webapp']], userRemoteConfigs: [[url: 'https://github.com/tomerschwartz24/Go-Projects']]])
            }
        }
        stage('Build golang website binary') {
            steps {
              sh 'go build website.go'
            }
        }
        stage('Build Website') {
            steps {
                sh 'docker build -t tomer/webapp -f Dockerfile .'
            }
        }
        
        stage('Push to Hub') {
            steps {
                withCredentials([usernamePassword( credentialsId: 'docker_hub', usernameVariable: 'USER', passwordVariable: 'PASSWORD')]) {
                sh 'docker build -t tomer/webapp -f Dockerfile .'
                }
            }
        }
    }
}