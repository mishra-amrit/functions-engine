# <img src="https://raw.githubusercontent.com/mishra-amrit/FunctionsEngine/master/assets/logo.png" width="50%" height="50%">

<b> Functions Engine lets you deploy your own Function as a Service / Serverless Platform.</b>

<b>1. Code > 2. Package as Jar or zip > 3. Upload to FunctionsEngine > 4. Ready to Go !</b>

It's free & open source. Feel free to reach out for questions, suggestions & contributions to mishra.amrit@outlook.com

No SDK, No extra dependencies for your Apps. You just need to write your methods / functions with signatures as specified below, then just package it as jar or zip and upload it to FunctionsEngine.

    public <ResponseDTO-Class> <method-name>()

    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams)

    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams, 
         HashMap<String, Object> pathParams)

    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams, 
         HashMap<String, Object> pathParams, 
         <RequestDTO-Class> payload)

It's open source and supports functions written in Java. 

You can build your own serverless compute platform with single server or a cluster easily. 
It's very easy to use for developers and very easy to manage for admins/operators.

Written in Go

Note: Currently it supports only Java apps, future releases will also support Node.js & Go. Java apps should be uploaded as a jar, for apps built using Node.js, Go should be uploaded as zip/archives.

# Build the project 
    
    cd $PROJECT_HOME/engine
    "./build.sh"
    Note : You need to move to that directory where ./build.sh exists and execute it there.

   During the execution he bin directory will be deleted and recreated.\
   The "bin" directory will have the following subdirectories:

    bin/configuration - All the *config.yaml files are stored here.
    bin/log           - All the logs are stored and maintained here.
    bin/data          - All the data files (i.e deployement info, etc) are stored here.
    bin/deployments   - All the jars which are uploaded by the user will be stored
                        here in a uniquely named directory.

# Running Functions Engine from build output

    cd $PROJECT_HOME/engine
    "./run.sh"
    Note : You need to move to that directory where ./run.sh exists and execute it there.
     
  The above will execute the "FunctionsEngine" executable file in the bin directory.
    
    $PROJECT_HOME is the root of the project 
    i.e /home/test/work/FunctionsEngine/engine
    
# Functions Engine's Java 8 Runtime
<a href="https://github.com/mishra-amrit/myappengine/tree/master/appengine-runtimes/java8-runtime">https://github.com/mishra-amrit/FunctionsEngine/tree/master/functionRuntimes/java8-runtime</a>

