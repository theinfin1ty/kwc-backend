import graphene
from graphene_django import DjangoObjectType
from quiz.models import Seasons, Episodes, Questions, UserAnswers
from quiz import services


class SeasonType(DjangoObjectType):
    class Meta:
        model = Seasons
        fields = ("id", "name", "description", "air_date")


class EpisodeType(DjangoObjectType):
    class Meta:
        model = Episodes
        fields = ("id", "name", "description", "air_date")


class QuestionType(DjangoObjectType):
    class Meta:
        model = Questions
        fields = ("id", "question", "episode", "answer")


class UserAnswersType(DjangoObjectType):
    class Meta:
        model = UserAnswers
        fields = ("id", "user", "question", "answer")


class Query(graphene.ObjectType):
    seasons = graphene.List(SeasonType)
    season = graphene.Field(SeasonType, id=graphene.Int())
    episodes = graphene.List(EpisodeType, season_id=graphene.Int())
    episode = graphene.Field(EpisodeType, id=graphene.Int())
    questions = graphene.List(QuestionType, episode_id=graphene.Int())
    question = graphene.Field(QuestionType, id=graphene.Int())
    user_answers = graphene.List(UserAnswersType, user_id=graphene.Int())
    user_answer = graphene.Field(
        UserAnswersType, user_id=graphene.Int(), question_id=graphene.Int()
    )

    def resolve_seasons(self, root, info):
        return services.get_all_seasons()

    def resolve_season(self, root, info, id):
        return services.get_season(id)

    def resolve_episodes(self, root, info, season_id):
        return services.get_all_episodes(season_id)

    def resolve_episode(self, root, info, id):
        return services.get_episode(id)

    def resolve_questions(self, root, info, episode_id):
        return services.get_all_questions(episode_id)

    def resolve_question(self, root, info, id):
        return services.get_question(id)

    def resolve_user_answers(self, root, info, user_id):
        return services.get_all_user_answers(user_id)

    def resolve_user_answer(self, root, info, user_id, question_id):
        return services.get_user_answer(user_id, question_id)


class Mutation(graphene.ObjectType):
    update_season = graphene.Field(
        SeasonType,
        description="Update a season",
        id=graphene.Int(),
        name=graphene.String(),
        season_description=graphene.String(),
        air_date=graphene.String(),
    )

    update_episode = graphene.Field(
        EpisodeType,
        description="Update a episode",
        id=graphene.Int(),
        name=graphene.String(),
        episode_description=graphene.String(),
        air_date=graphene.String(),
        season=graphene.Int(),
    )
    update_question = graphene.Field(
        QuestionType,
        description="Update a question",
        id=graphene.Int(),
        question=graphene.String(),
        episode=graphene.Int(),
        answer=graphene.String(),
    )
    update_user_answer = graphene.Field(
        UserAnswersType,
        description="Update a user's answer",
        id=graphene.Int(),
        user=graphene.Int(),
        question=graphene.Int(),
        answer=graphene.String(),
    )

    def resolve_update_season(self, root, info, id, name, season_description, air_date):
        return services.update_season(id, name, season_description, air_date)

    def resolve_update_episode(self, root, info, id, name, episode_description, air_date, season):
        return services.update_episode(id, name, episode_description, air_date, season)

    def resolve_update_question(self, root, info, id, question, episode, answer):
        return services.update_question(id, question, episode, answer)

    def resolve_update_user_answer(self, root, info, id, user, question, answer):
        return services.update_user_answer(id, user, question, answer)


schema = graphene.Schema(query=Query, mutation=Mutation)
