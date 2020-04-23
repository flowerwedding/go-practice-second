import requests
from lxml import etree
from pip._vendor.distlib.compat import raw_input  # 配置各种库，好多都是写了后点下划线然后上面自动生成

xuehao = raw_input('请输入学号：')
html = requests.get("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + xuehao)  # 跳转到该网页，教务在线的课表查询
# print(html.text)  # xx.text是该URL的源码

etree_html = etree.HTML(html.text)
nameURL = '//*[@id="head"]/div/div[1]/ul/li[2]/text()'  # fn+f12 >> 找目录 >> copy xpath
classURL = '//*[@id="kbStuTabs-list"]/div/table/tbody/tr/td/text()'


def demo(f, URL):
    global name, time, place, time2, time3  # 这行写完就存在了，可能系统对初学者太友好了
    content = etree_html.xpath(URL)  # 获取网页内容
    for each in content:  # 给每次获得的那行HTTP代码去掉标签及之外的东西
        replace = each.replace('\n', '').replace(' ', '')
        if replace == '-' or replace == '':
            continue
        else:  # 正文开始：
            if f == 0:  # 这个是给姓名的返回
                return replace
            else:  # 给课表信息的返回
                d = replace[:1]  # 截取获得的字符串的第一个
                if f == 1 and (d == 'A' or d == 'B'):  # 所有课程名的代号都是A或B开头，课程名由课程代号+课程名组成
                    name = '    ' + replace[9:]
                    f = 2
                e = replace[:1]  # 从f=1到f=4是一个大标签里的要名字（2行）时间（7行）地点（8行）
                if f == 2 and e == '星':
                    time = '--' + replace
                    f = 3
                if f == 4 and e == '星':  # 因为一周要多次课，但后面课的时间在大标签之后，不知道为什么反正试出来汉字占1字节
                    time3 = time2
                    time2 = '--' + replace
                    if time3 == time2: time3 = ''
                if f == 4 and (d == '理' or d == '实'):  # 课表信息的输出，同时f变回1开始下一个大标签的读取
                    f = 1
                    if time3 != '':
                        print(name + time + time2 + time3 + place)
                    elif time2 != '':
                        print(name + time + time2 + place)
                    else:
                        print(name + time + place)
                if f == 3 and e != '星':  # 地点在的第8行也就是最后一行，只要找到时间那下一行就是了
                    place = '--' + replace
                    f = 4
                    time2 = time3 = ''  # 我感觉一周课最多那门也就三次吧，不想写数组，虽然我知道array和list是正解


f = 0
name = demo(f, nameURL)  # 自定义函数去……，主要还是觉得一次次去空格太懒了
print(name[32:] + '同学,你本学期的课表为:')
f = 1
demo(f, classURL)  # 这里没return，没参数需要返回
