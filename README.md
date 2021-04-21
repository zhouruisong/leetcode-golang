# leetcode-golang

# 本地部署cpeh
    docker run -itd --restart=always --name ceph-mon --network ceph-network --ip 172.20.0.10 \
    -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_IP=172.20.0.10 -e MON_NAME=ceph-mon -e CEPH_PUBLIC_NETWORK=172.20.0.0/16 \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    -v /data/ceph/var/log/ceph/:/var/log/ceph/ \
    ceph/daemon mon

    搭建osd节点
    docker exec ceph-mon ceph auth get client.bootstrap-osd -o /var/lib/ceph/bootstrap-osd/ceph.keyring
    
    修改配置文件以兼容etx4硬盘
    sudo vi /etc/ceph/ceph.conf
    在文件最后添加：
    
    osd max object name len = 256
    osd max object namespace len = 64
    
    分别启动三个容器来模拟集群
    
    docker run -itd --restart=always --privileged=true --name ceph-osd-0 --network ceph-network --ip 172.20.0.11 \
    -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    -v /data/ceph/var/lib/ceph/osd/0:/var/lib/ceph/osd \
    -v /etc/localtime:/etc/localtime:ro \
    ceph/daemon osd
    
    docker run -itd --restart=always --privileged=true --name ceph-osd-1 --network ceph-network --ip 172.20.0.12 \
    -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    -v /data/ceph/var/lib/ceph/osd/1:/var/lib/ceph/osd \
    -v /etc/localtime:/etc/localtime:ro \
    ceph/daemon osd
    
    docker run -itd --restart=always --privileged=true --name ceph-osd-2 --network ceph-network --ip 172.20.0.13 \
    -e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    -v /data/ceph/var/lib/ceph/osd/2:/var/lib/ceph/osd \
    -v /etc/localtime:/etc/localtime:ro \
    ceph/daemon osd
    
    
    5. 搭建mgr节点
    
    docker run -itd --restart=always --privileged=true --name ceph-mgr --network ceph-network --ip 172.20.0.14 \
    -e CLUSTER=ceph -p 7000:7000 --pid=container:ceph-mon \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    ceph/daemon mgr
    
    开启管理界面
    docker exec ceph-mgr ceph mgr module enable dashboard
    
    docker exec ceph-mon ceph auth get client.bootstrap-rgw -o /var/lib/ceph/bootstrap-rgw/ceph.keyring
    
    docker run -itd --restart=always --privileged=true --name ceph-rgw --network ceph-network --ip 172.20.0.15 \
    -e CLUSTER=ceph -e RGW_NAME=ceph-rgw -p 7480:7480 \
    -v /data/ceph/var/lib/ceph/:/var/lib/ceph/ \
    -v /data/ceph/etc/ceph:/etc/ceph \
    -v /etc/localtime:/etc/localtime:ro \
    ceph/daemon rgw
    
    
    3.6.1 开启dashboard功能
    
    在master节点执行：
    docker exec mgr ceph mgr module enable dashboard
    
    设置用户名秘密
    vi adminpw 里面设置密码admin123
    //添加用户
    ceph dashboard ac-user-create admin -i adminpw administrator
    账号：admin 密码：admin123
    
    3.6.4 配置外部访问端口
    docker exec ceph-mgr ceph config set mgr mgr/dashboard/server_port 7000
    3.6.5 配置外部访问IP
    docker exec ceph-mgr ceph config set mgr mgr/dashboard/server_addr 0.0.0.0
    3.6.6 重启Mgr DashBoard服务
    docker restart ceph-mgr
    3.6.7 查看Mgr DashBoard服务信息
    
    docker exec ceph-mgr ceph mgr services
    -------------------------------------------
    {
       "dashboard": "http://10.10.85.38:7000"
    }
    
    //创建用户
    radosgw-admin user create --uid=test01 --display-name=test01 --access-key=test01 --secret-key=test01 --email=test01@qq.com
    
    //查询pool
    ceph osd lspools
    //查询对象在那些osd上
    osd map zrs myobject
    //查询指定pool下的对象
    ceph osd map zrs
    //往指定pool打入数据压侧
    rados -p zrs bench 20 write --no-cleanup

------------------------------------------
定时查看内存占用

    #!/bin/bash
    prog_name="chrome"
    prog_mem=$(pidstat -r -u -h -C $prog_name |awk 'NR==4{print $12}')
    time=$(date "+%Y/%m/%d/%H:%M:%S")
    echo $time" "$prog_mem >>/data/tmp/prog_mem.log

输出格式,放到excel中生成图表

    2021/04/21/10:12:19 1381252
    2021/04/21/10:12:20 1381252
    2021/04/21/10:12:21 1381292
    2021/04/21/10:12:22 1381182
   
excel生成图表

    ![image](https://github.com/zhouruisong/leetcode-golang/blob/master/excel.png)
