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
    processing = false;

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

    /**
     * Log user in the platform.
     * Shows validation errors if form submission is invalid.
     */
    login() {
        this.submitted = true;
        this.processing = true;

        // stop here if form is invalid
        if (this.loginForm.invalid) {
            this.processing = false;
            return
        }

        this._footballService.login(this.loginForm.value).then(response => {
            this.submitted = false;
            this.processing = false;

            if (response.success) {
                sessionStorage.setItem('username', this.loginForm.value.username);
                sessionStorage.setItem('userId', response.data.userId);

                this._footballService.changeMessage('true');
                this._footballService.changeUsernameSource(this.loginForm.value.username);
                // this._router.navigate(['/fixtures']);
                window.location.href = 'http://localhost:4200/fixtures';
            }
        }).catch(error => { 
            this.processing = false;
            this._flashMessageService.show('Unsuccessful login. Please try again.', {
                cssClass: 'alert-danger',
                timeout: 5000
            });
        })
    }
}