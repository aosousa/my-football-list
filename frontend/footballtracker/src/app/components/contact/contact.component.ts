import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'contact',
    templateUrl: './contact.component.html',
    styleUrls: ['./contact.component.css']
})
export class ContactComponent implements OnInit {
    contactForm: FormGroup;
    submitted = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService
    ) {
        this.contactForm = this._formBuilder.group({
            type: ['', Validators.required],
            subject: ['', Validators.required],
            message: ['', Validators.required]
        });
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Contact")
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.contactForm.controls;
    }

    submitContact() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.contactForm.invalid) {
            return
        }

        this._footballService.sendEmail(this.contactForm.value).then(response => {
            this.submitted = false;

            if (response.success) {
                this.contactForm.reset();
                this._flashMessagesService.show('Message sent successfully!', {
                    cssClass: 'alert-success',
                    timeout: 5000
                });
            }
        }).catch(error => {
            this._flashMessagesService.show('An error occurred while trying to submit the contact form. Please try again later.', {
                cssClass: 'alert-danger',
                timeout: 5000
            });
        });
    }
}