import collections
from importlib.abc import ResourceReader
from flask import Flask
from flask_restful import reqparse, abort, Api, Resource
# from flask_cors import CORS
from pymongo import MongoClient
import os

app = Flask(__name__)
api = Api(app)

parser = reqparse.RequestParser()

client = MongoClient('spotrate-db', 27017)
db = client['reviews']
collection = db['pitchfork']

# CORS(api)

class HelloWorld(Resource):
    def get(self):
        return {'hello': 'world'}

class Reviews(Resource):
    def get(self, project_title, artist):
        pitchforkReview = collection.find_one({ "title": project_title, "artist": artist})
        
        return pitchforkReview, 200
        
        
    def post(self):
        args = parser.parse_args()
        newReviewId = collection.insert_one(args).inserted_id
        
        return newReviewId, 201

api.add_resource(HelloWorld, '/')