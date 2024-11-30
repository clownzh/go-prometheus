pipeline {
    agent any

    environment {
        // 从 Jenkins 全局变量中读取 Docker 用户名和密码
        BRANCH_NAME = "${env.BRANCH_NAME}"  // 使用 Jenkins 提供的环境变量来获取分支名称
        IMAGE_NAME = "registry.cn-hangzhou.aliyuncs.com/hisyygh/go-prometheus"
        BUILD_DATE = sh(returnStdout: true, script: 'date +%Y%m%d%H%M%S').trim()
        IMAGE_TAG = "${IMAGE_NAME}:${BUILD_DATE}"
    }

    stages {
        stage('Checkout Code') {
            steps {
                // 从 GitHub 拉取代码，并切换到指定分支
                git url: 'https://github.com/clownzh/go-prometheus.git', branch: "${BRANCH_NAME}"
            }
        }

          stage('Build Docker Image') {
                    steps {
                        // 构建 Docker 镜像，标签为当前年月日时分秒
                        sh """
                            docker build -t ${IMAGE_TAG} .
                        """
                    }
                }


        stage('Login to Docker Registry') {
            steps {
                // 登录到阿里云 Docker 注册表
                 withCredentials([usernamePassword(credentialsId: 'AliRegistry', passwordVariable: 'AliRegistryPassword', usernameVariable: 'AliRegistryUser')]) {
                        sh "docker login -u ${AliRegistryUser} registry.cn-hangzhou.aliyuncs.com -p ${AliRegistryPassword}"
                    }
            }
        }

        stage('Push Docker Image') {
            steps {
                // 推送 Docker 镜像到阿里云注册表
                sh """
                    docker push ${IMAGE_TAG}
                """
            }
        }
    }

    post {
        always {
            // 清理工作空间
            cleanWs()
        }
        success {
            // 如果构建成功，输出成功信息
            echo "Docker image pushed successfully: ${IMAGE_TAG}"
        }
        failure {
            // 如果构建失败，输出失败信息
            echo "Docker image push failed"
        }
    }
}