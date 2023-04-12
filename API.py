from flask import Flask, request

app = Flask(__name__)

code = None
@app.route("/postCode", methods=["POST"])
def postCode():

    input_json = request.get_json(force=True)

    print("Got data from client: ", input_json)
    code = input_json
    return "Code received"


@app.route("/getCode", methods=["GET"])
def getCode():

    if code is None:
        return "{\"code\": \"No code\"}"
    return code


app.run(port=3000)