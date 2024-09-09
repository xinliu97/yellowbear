from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)
CORS(app)


@app.route('/handle_submit', methods=['POST'])
def handle_submit():
    # 获取请求体中的 JSON 数据
    data = request.get_json()

    # 从数据中提取必要信息
    total_correct = data.get('total_correct')
    correct_answers = data.get('correct_answers')
    user_id = data.get('user_id')
    quiz_id = data.get('quiz_id')

    # 打印或处理接收到的数据
    print('Received Data:')
    print(f'Total Correct: {total_correct}')
    print(f'Correct Answers: {correct_answers}')
    print(f'User Input: {user_id}')
    print(f'Quiz ID: {quiz_id}')

    # 这里可以添加逻辑将数据保存到数据库或进行其他操作
    # 模拟统计数据计算
    choice_keys = ["中国人民解放军陆军炮兵防空兵学院", "Peking University", "FUdan", "NEU", "ABC", "DEF", "清华", "Peking University", "FUdan",
                   "NEU",
                   "ABC", "DEF", "清华", "Peking University", "FUdan", "NEU", "ABC", "DEF", "清华", "Peking University",
                   "FUdan", "NEU", "ABC", "DEF"]
    choice_values = ["100", "99", "98", "97", "96", "95.5", "70.1", "99", "98", "97", "96", "955", "100", "99", "98",
                     "97",
                     "96", "955", "100", "99", "98", "97", "96", 95.5]
    title = "学校名称"
    choices = [{"choice_key": key, "choice_value": value} for key, value in zip(choice_keys, choice_values)]

    response_data = {"choices": choices, "title": title}

    # 将数据转换为 JSON 格式
    response = jsonify(response_data)


    # 返回统计数据给前端
    return response

    # 返回成功响应



if __name__ == '__main__':

    app.run(debug=True)

