import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
    loginForm: FormGroup;
    submitted = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _router: Router,
    ) {
        this.loginForm = this._formBuilder.group({
            username: ['', Validators.required],
            password: ['', Validators.required]
        });
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Login")
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.loginForm.controls;
    }

    login() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.loginForm.invalid) {
            return
        }

        this._footballService.login(this.loginForm.value).then(response => {
            this.submitted = false;
            if (response.success) {
                // TEMPORARY: redirect to /fixtures when that route exists
                this._footballService.changeMessage('true');
                this._footballService.changeUsernameSource(this.loginForm.value.username);
                this._router.navigate(['/']);
            } else {
                this._flashMessageService.show('Unsuccessful login. Please try again.', {
                    cssClass: 'alert-danger',
                    timeout: 5000
                });
            }
        })
    }
}