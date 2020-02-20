from flask import Flask, render_template, flash, request, abort, redirect
from db import *

app = Flask(__name__)


# 登录
@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == 'POST':
        username = request.args['username']
        password = request.args['password']
        # print("username:", username, "password:", password)
        if isExisted(username, password):
            return "login successful"
        else:
            message = 'Login Failed'
            return render_template('index.html', message=message)
    return render_template('index.html')


# 注册
@app.route("/register", methods=["GET", "POST"])
def register():
    if request.method == 'POST':
        username = request.args['username']
        password = request.args['password']
        # print("username:", username, "password:", password)
        addUser(username, password)
        return "Registered successfully"
    return render_template("index.html")


if __name__ == '__main__':
    app.run()
