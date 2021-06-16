import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { InjectSetupWrapper } from '@angular/core/testing';
import { Observable } from "rxjs";
import { Todo } from './todo';
import { Note } from './todo';

@Injectable()

export class TodoService {

  constructor(private http: HttpClient) {}

  getTodo() {
    return this.http.get<Todo>('http://localhost:8080');
  }

  addTodo(nt: Note) {
    console.log(nt)
    return this.http.post<Note>('http://localhost:8080/addNote', nt);
  }
}
