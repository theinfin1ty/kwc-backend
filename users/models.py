from django.db import models
from django.contrib.auth.models import AbstractBaseUser

# Create your models here.
class Users(AbstractBaseUser):
  name = models.CharField(max_length=255, blank=False)
  email = models.CharField(max_length=255, blank=False, unique=True)
  is_active = models.BooleanField(default=True)
  username = models.CharField(max_length=255, blank=True, unique=False)

  USERNAME_FIELD = 'email'
  EMAIL_FIELD = 'email'

  def __repr__(self):
      return self.email