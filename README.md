# myAppEngine - Serverless made simple
 
<b>1. Code > 2. Package as Jar or zip > 3. Upload to MyAppEngine > 4. Ready to Go !</b>

No SDK, No extra dependencies for your Apps. You just need to write your methods / functions with signatures as specified below. Then just package it as jar or zip and upload it to MyAppEngine.

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

Note: Currently it supports only Java apps, future releases will also support Node.js & Go.
Java apps should be uploaded as a jar, for apps built using Node.js, Go should be uploaded as zip/archives.

<b>Build the project :</b>
    
    cd $PROJECT_HOME
    "./build.sh"
   
   During the execution he bin directory will be deleted and recreated.\
   The "bin" directory will have the following subdirectories:

    bin/configuration - All the *config.json files are stored here.
    bin/data          - All the data files (i.e deployement info, etc) are stored here.
    bin/deployments   - All the jars which are uploaded by the user will be stored
                        here in a uniquely named directory.

<b>Executable output from bin:</b>

    cd $PROJECT_HOME
    "./run.sh"
     
  The above will execute the "myappengine" executable file in the bin directory.
    
    $PROJECT_HOME is the root of the project 
    i.e /home/test/work/myappengine
