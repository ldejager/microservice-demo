import hudson.model.*;
import jenkins.model.*;

println "--> setting master executors to 2"
Jenkins.instance.setNumExecutors(2)
