import graphene
from graphene_django import DjangoObjectType
from quiz.models import Seasons, Episodes, Questions, UserAnswers

class SeasonType(DjangoObjectType):
    class Meta:
        model = Seasons
        fields = ("id", "name", "description", 'air_date')

class EpisodeType(DjangoObjectType):
    class Meta:
        model = Episodes
        fields = ("id", "name", "description", 'air_date')

class QuestionType(DjangoObjectType):
    class Meta:
        model = Questions
        fields = ("id", "question", "episode", "answer")

class UserAnswers(DjangoObjectType):
    class Meta:
        model = UserAnswers
        fields = ("id", "user", "question", "answer")

class Query(graphene.ObjectType):
    seasons = graphene.List(SeasonType)
    season = graphene.Field(SeasonType, id=graphene.Int())
    episodes = graphene.List(EpisodeType)
    episode = graphene.Field(EpisodeType, id=graphene.Int())
    questions = graphene.List(QuestionType)
    question = graphene.Field(QuestionType, id=graphene.Int())
    user_answers = graphene.List(UserAnswers)
    user_answer = graphene.Field(UserAnswers, id=graphene.Int())

    def resolve_seasons(root, info):
        return
    def resolve_season(root, info):
        return
    def resolve_episodes(root, info):
        return
    def resolve_episode(root, info):
        return
    def resolve_questions(root, info):
        return
    def resolve_question(root, info):
        return
    def resolve_user_answers(root, info):
        return
    def resolve_user_answer(root, info):
        return
    
class Mutation(graphene.ObjectType):
    update_season = graphene.Field(SeasonType, id=graphene.Int(), name=graphene.String())

schema = graphene.Schema(query=Query, mutation=Mutation)