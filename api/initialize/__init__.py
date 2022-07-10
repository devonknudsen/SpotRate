from flask import Flask
from flask_restful import Resource, Api
from flask_cors import CORS
import os

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = 'postgresql://postgres:postgres@spotrate-db:5432/spotrate?connection_limit=1'
app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
api = Api(app)
CORS(api)

class HelloWorld(Resource):
    def get(self):
        return {'hello': 'world'}

api.add_resource(HelloWorld, '/')