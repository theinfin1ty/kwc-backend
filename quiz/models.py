from django.db import models
from users.models import Users

# Create your models here.
class Seasons(models.Model):
    name = models.CharField(max_length=255)
    air_date = models.DateField(blank=True, unique=False)
    description = models.TextField(blank=True, unique=False)
  
    def __str__(self):
        return self.season_name

class Episodes(models.Model):
    name = models.CharField(max_length=255)
    air_date = models.DateField(blank=True, unique=False)
    description = models.TextField(blank=True, unique=False)
    season = models.ForeignKey(Seasons, on_delete=models.CASCADE)
    def __str__(self):
        return self.episode_name

class Questions(models.Model):
    question = models.CharField(max_length=255)
    answer = models.CharField(max_length=255)
    episode = models.ForeignKey(Episodes, on_delete=models.CASCADE)
    def __str__(self):
        return self.question
    
class UserAnswers(models.Model):
    user = models.ForeignKey(Users, on_delete=models.CASCADE)
    question = models.ForeignKey(Questions, on_delete=models.CASCADE)
    answer = models.CharField(max_length=255)
    def __str__(self):
        return self.user