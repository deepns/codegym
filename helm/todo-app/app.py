from flask import Flask, jsonify, request

app = Flask(__name__)

# In-memory data store
todos = []

# Get all todos
@app.route('/todos', methods=['GET'])
def get_todos():
    return jsonify(todos)

# Add a new todo
@app.route('/todos', methods=['POST'])
def add_todo():
    data = request.get_json()
    todo = data.get('todo')
    if todo:
        todos.append(todo)
        return jsonify({"message": "Todo added", "todos": todos}), 201
    else:
        return jsonify({"message": "Invalid input"}), 400

# Delete a todo by index
@app.route('/todos/<int:index>', methods=['DELETE'])
def delete_todo(index):
    if 0 <= index < len(todos):
        todos.pop(index)
        return jsonify({"message": "Todo deleted", "todos": todos})
    else:
        return jsonify({"message": "Invalid index"}), 400

# Health check endpoint
@app.route('/health', methods=['GET'])
def health_check():
    return jsonify({"status": "healthy"}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000)