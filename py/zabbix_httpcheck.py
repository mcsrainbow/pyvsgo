# Description: HTTP requests check for Zabbix
# Author: Dong Guo

# 文件描述: Zabbix HTTP 请求检查脚本 
import sys  # 导入系统模块
import requests  # 导入 requests 模块, 用于发送 HTTP 请求
import time  # 导入时间模块
import argparse  # 导入命令行解析模块
import textwrap  # 导入文本包装模块

def parse_opts():  # 定义参数解析函数
    """帮助消息(-h, --help)"""
    parser = argparse.ArgumentParser(  # 创建 ArgumentParser 对象
        formatter_class=argparse.RawDescriptionHelpFormatter,  # 设置描述格式器
        description=textwrap.dedent(  # 设置描述文本
        '''
        示例用法:
          {0} -u idc1-web1/health  # 使用 URL 进行健康检查
          {0} -u http://idc1-web1/health  # 使用完整的 HTTP URL 进行健康检查
          {0} -u http://idc1-web1/health -c ok  # 在响应内容中查找特定字符串
          {0} -u http://idc1-web1/health -c ok -V  # 返回实际值而不是 0 和 1
          {0} -u http://idc1-web1/health -c ok -t 2 -V  # 设置超时时间和返回实际值
          {0} -u http://idc1-web2:3000  # 使用非标准端口进行健康检查
          {0} -u http://idc1-web3/login.php?page=redirect_string -a username:password -V  # 使用基本认证和返回实际值
          {0} -u https://idc2-web1.yourdomain.com -V  # 使用 HTTPS 和返回实际值
        '''.format(__file__)  # 格式化文件名
        ))

    parser.add_argument('-u', metavar='url', type=str, required=True, help='URL to GET or POST [default: http://]')  # 添加 URL 参数
    parser.add_argument('-t', metavar='timeout', type=float, help='seconds before connection times out [default: 10]')  # 添加超时时间参数
    parser.add_argument('-c', metavar='content', type=str, help='string to expect in the content')  # 添加响应内容参数
    parser.add_argument('-a', metavar='auth', type=str, help='username:password on sites with basic authentication')  # 添加基本认证参数
    parser.add_argument('-V', action="store_true", default=False, help='return actual value instead of 0 and 1')  # 添加返回实际值参数
    parser.add_argument('-p', metavar='payload', type=str, help='URL encoded http POST data')  # 添加 POST 数据参数

    if len(sys.argv) < 2:  # 如果命令行参数不足
        parser.print_help()  # 打印帮助信息
        sys.exit(2)  # 退出程序

    args = parser.parse_args()  # 解析命令行参数

    if args.a:  # 如果有基本认证参数
        if ':' not in args.a or len(args.a.split(':')) != 2:  # 检查格式是否正确
            print("Invalid auth format. Expected username:password")  # 打印错误信息
            sys.exit(2)  # 退出程序

    return {'url': args.u, 'timeout': args.t, 'content': args.c, 'auth': args.a, 'value': args.V, 'payload': args.p}  # 返回参数字典

def get_results(opts):  # 定义结果获取函数
    """使用给定参数获取结果"""
    url = opts['url']  # 获取 URL 参数
    if "http://" not in url and "https://" not in url:  # 如果 URL 不是完整的 HTTP 或 HTTPS 地址
        url = "http://" + url  # 添加 http:// 前缀

    start_timestamp = time.time()  # 记录开始时间
    if opts.get('timeout'):  # 如果有超时时间参数
        timeout = opts['timeout']  # 获取超时时间
    else:
        timeout = 10  # 默认超时时间为 10 秒

    try:  # 尝试发送 HTTP 请求
        if opts.get('auth'):  # 如果有基本认证参数
            from requests.auth import HTTPBasicAuth  # 导入基本认证模块
            username, password = opts['auth'].split(':')  # 拆分用户名和密码
            httpauth = HTTPBasicAuth(username, password)  # 创建基本认证对象
            if opts.get('payload'):  # 如果有 POST 数据参数
                payload = opts['payload']  # 获取 POST 数据
                req = requests.post(url, data=payload, auth=httpauth, timeout=timeout)  # 发送 POST 请求
            else:
                req = requests.get(url, auth=httpauth, timeout=timeout)  # 发送 GET 请求
        else:  # 如果没有基本认证参数
            if opts.get('payload'):  # 如果有 POST 数据参数
                payload = opts['payload']  # 获取 POST 数据
                req = requests.post(url, data=payload, timeout=timeout)  # 发送 POST 请求
            else:
                req = requests.get(url, timeout=timeout)  # 发送 GET 请求

        end_timestamp = time.time()  # 记录结束时间
        response_secs = round(end_timestamp - start_timestamp, 3)  # 计算响应时间

        if opts.get('value'):  # 如果要返回实际值
            if opts.get('content'):  # 如果有响应内容参数
                print(req.content.decode('utf-8'))  # 打印响应内容
            elif opts.get('timeout'):  # 如果有超时时间参数
                print(response_secs)  # 打印响应时间
            else:
                print(req.status_code)  # 打印状态码
        else:  # 如果不返回实际值
            if req.status_code == requests.codes.ok:  # 如果状态码是 OK
                if opts.get('content'):  # 如果有响应内容参数
                    if opts['content'] in req.content.decode('utf-8'):  # 检查是否包含特定字符串
                        print(0)  # 打印 0 表示成功
                    else:
                        print(1)  # 打印 1 表示失败
                else:
                    print(0)  # 打印 0 表示成功
            else:  # 如果状态码不是 OK
                print(1)  # 打印 1 表示失败

    except requests.exceptions.Timeout:  # 捕获超时异常
        print("Timeout" if opts.get('value') else 1)  # 打印超时信息或 1

    except requests.exceptions.ConnectionError:  # 捕获连接错误异常
        print("ConnectionError" if opts.get('value') else 1)  # 打印连接错误信息或 1

    except Exception as e:  # 捕获其他异常
        print(f"Unexpected error: {str(e)}" if opts.get('value') else 1)  # 打印异常信息或 1
        return 2  # 返回 2 表示失败

    return 0  # 返回 0 表示成功

def main():  # 定义主函数
    opts = parse_opts()  # 解析命令行参数
    return get_results(opts)  # 获取结果

if __name__ == '__main__':  # 如果是主模块
    sys.exit(main())  # 运行主函数并退出程序