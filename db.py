import pymysql

# 创建数据库连接，本地数据库，用户名为root，密码，数据库为dome7
conn = pymysql.connect("localhost", "root", "", "dome7", )
# 创建游标
cur = conn.cursor()


def addUser(username, password):
    # 将username和password插入到数据库中
    sql = "insert into test (name, password) values ('%s', '%s')" % (username, password)
    # 执行方法
    cur.execute(sql)
    # 保持
    conn.commit()
    # 关闭连接
    conn.close()


def isExisted(username, password):
    sql = "select * from test where name='%s' and password = '%s'" % (username, password)
    cur.execute(sql)
    result = cur.fetchall()
    if len(result) == 0:
        return False
    else:
        return True
