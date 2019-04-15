import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

@Component({
	selector: 'index',
	templateUrl: './index.component.html',
	styleUrls: ['./index.component.css']
})
export class IndexComponent implements OnInit {
	constructor(
		private _titleService: Title
	) { }

	ngOnInit() {
		this._titleService.setTitle("Football Tracker")
	}
}
