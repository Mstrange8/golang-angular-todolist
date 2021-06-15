import { Component } from '@angular/core';
import { OnInit } from '@angular/core';

import { TodoService } from "./todo/todo.service";
import { Todo } from "./todo/todo";
import { Note } from "./todo/todo";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  constructor(private td: TodoService) {}

  todo: Todo;
  note: Note;

  ngOnInit() {
    this.td.getTodo()
    .subscribe((todo) => {
      this.todo = todo;
    });
  }

  add() {
    this.td.addTodo(this.note)
    .subscribe((note) => {
      this.note.title = note.title,
      this.note.description = note.description
    })
  }
}
