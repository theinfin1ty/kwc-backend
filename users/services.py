from users.models import Users

def get_all_users():
  return Users.objects.all()

def get_user(id):
  return Users.objects.get(id=id)

def update_user(id, name):
  user = Users.objects.get(id=id)
  user.name = name
  user.save()
  return user