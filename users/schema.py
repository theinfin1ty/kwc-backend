import graphene
from graphene_django import DjangoObjectType
from users.models import Users
from users import services

class UserType(DjangoObjectType):
    class Meta:
        model = Users
        fields = ("id", "name", "email")

class Query(graphene.ObjectType):
    users = graphene.List(UserType)
    user = graphene.Field(UserType, id=graphene.Int())

    def resolve_users(root, info):
        return services.get_all_users()
    def resolver_user(root, info, id):
        return services.get_user(id)

class Mutation(graphene.ObjectType):
    update_user = graphene.Field(UserType, id=graphene.Int(), name=graphene.String())

    def resolve_update_user(root, info, id, name):
        return services.update_user(id, name)

schema = graphene.Schema(query=Query, mutation=Mutation)