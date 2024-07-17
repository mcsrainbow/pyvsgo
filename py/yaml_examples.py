import json
import logging
import yaml
import os

# 配置日志记录器
logging.basicConfig(level=logging.INFO, format='%(asctime)s %(levelname)s: %(message)s', datefmt='%Y/%m/%d %H:%M:%S')
info_logger = logging.getLogger("info_logger")
error_logger = logging.getLogger("error_logger")

def log_with_level(logger, level, msg, *args):
    if level == 'INFO':
        logger.info(msg, *args)
    elif level == 'ERROR':
        logger.error(msg, *args)

def main():
    tmp_file = "tmp/yaml_examples.yaml"

    # 生成 YAML 文件的数据
    data = {
        "key1": "value1",
        "key2": {
            "subkey1": "subvalue1",
            "subkey2": "subvalue2"
        },
        "key3": [1, 2, 3]
    }

    log_with_level(info_logger, "INFO", "生成 YAML 的初始数据:")
    try:
        out = json.dumps(data, indent=2)
        print(out)
    except Exception as e:
        log_with_level(error_logger, "ERROR", "JSON 序列化失败: %s", str(e))
        return

    # 将数据写入 YAML 文件
    try:
        with open(tmp_file, 'w') as file:
            yaml.dump(data, file)
    except PermissionError as e:
        log_with_level(error_logger, "ERROR", "写入 YAML 文件失败: %s", str(e))
        return
    except Exception as e:
        log_with_level(error_logger, "ERROR", "YAML 序列化失败: %s", str(e))
        return

    # 更新 key2.subkey1 数据
    data["key2"]["subkey1"] = "new_subvalue1"

    # 将更新后的数据写回 YAML 文件
    try:
        with open(tmp_file, 'w') as file:
            yaml.dump(data, file)
    except PermissionError as e:
        log_with_level(error_logger, "ERROR", "写入 YAML 文件失败: %s", str(e))
        return
    except Exception as e:
        log_with_level(error_logger, "ERROR", "YAML 序列化失败: %s", str(e))
        return
    log_with_level(info_logger, "INFO", "更新 YAML key2.subkey1: %s", data["key2"]["subkey1"])

    # 读取 YAML 文件以验证更新
    try:
        with open(tmp_file, 'r') as file:
            updated_data = yaml.safe_load(file)
    except PermissionError as e:
        log_with_level(error_logger, "ERROR", "读取 YAML 文件失败: %s", str(e))
        return
    except Exception as e:
        log_with_level(error_logger, "ERROR", "解析 YAML 文件失败: %s", str(e))
        return

    log_with_level(info_logger, "INFO", "读取更新后的 YAML 文件内容 'tmp/yaml_examples.yaml':")
    try:
        out = yaml.dump(updated_data, default_flow_style=False)
        print(out)
    except Exception as e:
        log_with_level(error_logger, "ERROR", "YAML 序列化失败: %s", str(e))
        return

    # 校验更新是否成功
    if updated_data["key2"]["subkey1"] != "new_subvalue1":
        log_with_level(error_logger, "ERROR", "更新失败")
        return
    log_with_level(info_logger, "INFO", "更新校验成功")

if __name__ == "__main__":
    if not os.path.exists('tmp'):
        os.makedirs('tmp')
    main()

