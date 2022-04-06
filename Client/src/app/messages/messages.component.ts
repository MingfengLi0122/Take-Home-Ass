import { Component, OnInit } from '@angular/core';
import { MessageService } from '../message.service';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})

export class MessagesComponent implements OnInit {
  // must be public becz we want to bind to this template
  constructor(public messageService: MessageService) { }

  ngOnInit(): void {
  }

}
