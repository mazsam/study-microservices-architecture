
"""The Python implementation of the GRPC helloworld.Greeter server."""

from concurrent import futures

import logging
import grpc
import sqlalchemy as db

from typing import Optional
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped
from sqlalchemy.orm import mapped_column

protos, services = grpc.protos_and_services("user.proto")

# Create test.sqlite automatically
engine = db.create_engine('sqlite:///user.sqlite')
connection = engine.connect()
metadata = db.MetaData()

user = db.Table('user', metadata,
               db.Column('user_id', db.String(100), nullable=False),
               db.Column('first_name', db.String(255), nullable=True),
               db.Column('last_name', db.String(255), nullable=True),
               db.Column('email', db.String(255), nullable=True),
               db.Column('phone_number', db.String(255), nullable=True),

               )
metadata.create_all(engine)


class User(services.UserServicer):
    def CreateProfile(self, request, context):
        engine = db.create_engine('sqlite:///user.sqlite')
        connection = engine.connect()
        
        query = db.insert(user).values(user_id=request.user_id, first_name=request.first_name, last_name=request.last_name, email=request.email, phone_number=request.phone_number) 

        ResultProxy = connection.execute(query)

        return protos.BaseReply(message='Hello, %s!' % request.user_id, code="0")

    def UpdateProfile(self, request, context):
        query = db.update(user).values(first_name=request.first_name, last_name=request.last_name, email=request.email, phone_number=request.phone_number) 
        query = query.where(user.columns.user_id==request.user_id)

        ResultProxy = connection.execute(query)
        return protos.BaseReply(message='Hello, %s!' % request.user_id, code="0")

    def GetProfile(self, request, context):
        query = db.where(user)
        query = query.where(user.columns.user_id==request.user_id)

        connection.execute(query)
        return protos.BaseReply(message='Hello, %s!' % request.user_id, code="0")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    services.add_UserServicer_to_server(User(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
