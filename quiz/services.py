from users.models import Users
from quiz.models import Seasons, Episodes, Questions, UserAnswers


def get_all_seasons():
    """
    Retrieve all seasons from the database.

    Returns:
        QuerySet: A QuerySet containing all Seasons objects.
    """
    return Seasons.objects.all()


def get_season(season_id):
    """
    Retrieve a specific season by its ID from the database.

    Args:
        season_id (int): The ID of the season to retrieve.

    Returns:
        Season: The Season object with the specified ID.
    """
    return Seasons.objects.get(id=season_id)


def update_season(season_id, name, description, air_date):
    """
    Update a season's information in the database.

    Args:
        season_id (int): The ID of the season to update.
        name (str): The new name of the season.
        description (str): The new description of the season.
        air_date (datetime): The new air date of the season.

    Returns:
        Season: The updated Season object.
    """
    season = Seasons.objects.get(id=season_id)
    season.name = name
    season.description = description
    season.air_date = air_date
    season.save()
    return season


def get_all_episodes(season_id):
    """
    Retrieve all episodes for a specific season from the database.

    Args:
        season_id (int): The ID of the season for which to retrieve episodes.

    Returns:
        QuerySet: A QuerySet containing all Episodes objects for the specified season.
    """
    return Episodes.objects.filter(season=season_id)


def get_episode(episode_id):
    """
    Retrieve a specific episode by its ID from the database.

    Args:
        episode_id (int): The ID of the episode to retrieve.

    Returns:
        Episode: The Episode object with the specified ID.
    """
    return Episodes.objects.get(id=episode_id)


def update_episode(episode_id, name, description, air_date, season):
    """
    Update an episode's information in the database.

    Args:
        episode_id (int): The ID of the episode to update.
        name (str): The new name of the episode.
        description (str): The new description of the episode.
        air_date (datetime): The new air date of the episode.
        season (Season): The season to which the episode belongs.

    Returns:
        Episode: The updated Episode object.
    """
    episode = Episodes.objects.get(id=episode_id)
    episode.name = name
    episode.description = description
    episode.air_date = air_date
    episode.season = season
    episode.save()
    return episode


def get_all_questions(episode_id):
    """
    Retrieve all questions for a specific episode from the database.

    Args:
        episode_id (int): The ID of the episode for which to retrieve questions.

    Returns:
        QuerySet: A QuerySet containing all Questions objects for the specified episode.
    """
    return Questions.objects.filter(episode=episode_id)


def get_question(question_id):
    """
    Retrieve a specific question by its ID from the database.

    Args:
        id (int): The ID of the question to retrieve.

    Returns:
        Question: The Question object with the specified ID.
    """
    return Questions.objects.get(id=question_id)


def update_question(question_id, question, episode, answer):
    """
    Update a question's information in the database.

    Args:
        question_id (int): The ID of the question to update.
        question (str): The new question text.
        answer (str): The new answer to the question.
        episode (Episode): The episode to which the question belongs.

    Returns:
        Question: The updated Question object.
    """
    question = Questions.objects.get(id=question_id)
    question.question = question
    question.answer = answer
    question.episode = episode
    question.save()
    return question


def get_all_user_answers(user_id):
    """
    Retrieve all user answers for a specific user from the database.

    Args:
        user_id (int): The ID of the user for whom to retrieve answers.

    Returns:
        QuerySet: A QuerySet containing all UserAnswers objects for the specified user.
    """
    return UserAnswers.objects.filter(user=user_id)


def get_user_answer(user_id, question_id):
    """
    Retrieve a specific user answer by its ID from the database.

    Args:
        id (int): The ID of the user answer to retrieve.

    Returns:
        UserAnswer: The UserAnswer object with the specified ID.
    """
    return UserAnswers.objects.get(user=user_id, question=question_id)


def update_user_answer(user_id, user, question, answer):
    """
    Update a user's answer to a question in the database.

    Args:
        id (int): The ID of the user answer to update.
        user (User): The user who provided the answer.
        question (Question): The question to which the answer belongs.
        answer (str): The updated answer text.

    Returns:
        UserAnswer: The updated UserAnswer object.
    """
    user_answer = UserAnswers.objects.get(id=user_id)
    user_answer.user = user
    user_answer.question = question
    user_answer.answer = answer
    user_answer.save()
    return user_answer
