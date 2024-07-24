from django.urls import path, re_path
from . import views

urlpatterns = [
    re_path(r'^api/show_books/?$', views.show_books, name='show_books'),
    re_path(r'^api/add_book/?$', views.add_book, name='add_book'),
    re_path(r'^api/delete_book/(?P<pk>\d+)/?$', views.delete_book, name='delete_book'),
    re_path(r'^api/update_book/(?P<pk>\d+)/?$', views.update_book, name='update_book'),
]
