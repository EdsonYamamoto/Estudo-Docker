from flask import Flask
from services.servico1 import classeServico1

app = Flask(__name__)
classe = classeServico1()

@app.route('/', methods=['GET'])
def hello_world():
	classe.primeiraFuncao()
	return 'executado funcao de uma classe singleton'

@app.route('/v1', methods=['GET'])
def segunda_world():
	return classe.segundaFuncao()

if __name__ == '__main__':
	app.run(debug=True, host='0.0.0.0')