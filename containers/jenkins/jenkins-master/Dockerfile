FROM rwxlabs/jenkins-volumes

COPY src/plugins.sh /usr/local/bin/install_plugins.sh
COPY src/plugins.txt /usr/share/jenkins/plugins.txt

RUN /usr/local/bin/install_plugins.sh /usr/share/jenkins/plugins.txt

COPY src/master-executors.groovy /usr/share/jenkins/ref/init.groovy.d/
