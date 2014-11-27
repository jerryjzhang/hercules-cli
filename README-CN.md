##create

创建应用

    用法：hercules create <appName>
    示例:
        $ hercules create dsf
        Created application with id 1417019037871
    
##delete

删除应用

    用法: hercules delete <appName>
    示例:
        $ hercules delete dsf
        Deleted application dsf
    
##deploy

部署应用，支持三种部署方式：1. SVN源代码 2. docker镜像 3. dockerfile

    用法: hercules deploy <appName> (-s <svnURL> | -i <imageURI> | -f <filePath>) [-e <env>=<value>...]
    选项:
        -s, --svn-url <svnURL>  源代码的SVN地址
        -i, --docker-image <imageURI>  docker镜像的URI
        -f, --docker-file <filePath>  dockerfile的路径
        -e, --env <env>=<value>  应用级别的环境变量设置
    示例:
    $ hercules deploy dsf -s http://svnURL -e DB_URL=localhost:3006 
    Created release with id 141701903990

##scale

启动/扩缩容应用，可以单独为应用中每个程序设置进程(replica)个数。

    用法: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    示例:
        $ hercules scale dsf web=2 db=1
        Scaled application with 4 processes

##stop 

停止应用，该命令将停止应用所有的进程，要单独停止某个或某些进程，请使用process命令

    用法: hercules stop <appName>
    示例:
        $ hercules stop dsf
        Stopped application with 4 processes

##program

程序的管理命令，可以对程序增删改查

    用法: hercules program <appName> save -n <name> [-e <env>=<value>...  -c <programName>=<cmd>...  -p <programName>=<entrypoint>...]
          hercules program <appName> delete -n <name>
          hercules program <appName> list
    选项：
        -e, --env <env>=<val>  程序级别的环境变量设置
        -c, --cmd <programName>=<cmd>  程序的启动命令设置
        -p, --entrypoint <programName>=<entrypoint>  程序的入口命令设置
    示例:
        $ hercules program dsf save -n web -c web=start.sh -e DB_URL=localhost:3006 
        Saved program web

##release

发布包的管理命令，可以查询和回滚发布包
    
    用法: hercules release <appName> list 
          hercules release <appName> rollback -v <releaseVersion>
    选项:
        -v, --version <releaseVersion> 发布包的版本
    示例:
        $ hercules release dsf list
        ID              Version   Tag
        141701903990      v1      1.0.0 Release
        141701903991      v2      1.0.1 Release

##process

进程的管理命令，可以查询和停止进程

    用法：hercules process <appName> list
          hercules process <appName> stop -p <processID>
    选项:
        -p, --process-id <processID>  进程的ID
    示例:
        $ hercules process dsf list
         ID              Program   
        141701903990      web   
        141701903991      web 
        141701903992      db     
