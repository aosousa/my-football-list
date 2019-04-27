import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { AppRoutingModule } from '@app/app-routing.module';
import { HttpModule } from '@angular/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { environment } from 'environments/environment';

// 3rd party
import { FlashMessagesModule } from 'angular2-flash-messages';

// Pipes
import { CallbackPipe } from '@app/callback.pipe';

// Components
import { AboutComponent } from '@components/about/about.component';
import { AppComponent } from './app.component';
import { ContactComponent } from '@components/contact/contact.component';
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { FooterComponent } from '@components/footer/footer.component';
import { IndexComponent } from '@components/index/index.component';
import { LoginComponent } from '@components/login/login.component';
import { NavbarComponent } from '@components/navbar/navbar.component';
import { PrivacyPolicyComponent } from '@components/privacy-policy/privacy-policy.component';
import { ProfileComponent } from '@components/profile/profile.component';
import { ResetPasswordComponent } from '@components/reset-password/step1/reset-password.component';
import { NewPasswordComponent } from '@components/reset-password/step2/new-password.component';
import { SignupComponent } from '@components/signup/signup.component';
import { TermsOfServiceComponent } from '@app/components/terms-of-service/terms-of-service.component';

// Services
import { AuthGuard } from '@services/auth.guard';
import { ConfigService } from '@services/config.service';
import { FootballService } from '@services/football.service';
import { UtilsService } from '@services/utils.service';

export function ConfigLoader(configService: ConfigService) {
	return () => configService.loadConfig(environment.httpConfig)
}

@NgModule({
	declarations: [
		AboutComponent,
		AppComponent,
		CallbackPipe,
		ContactComponent,
		FixturesComponent,
		FooterComponent,
		IndexComponent,
		LoginComponent,
		NavbarComponent,
		PrivacyPolicyComponent,
		ProfileComponent,
		ResetPasswordComponent,
		NewPasswordComponent,
		SignupComponent,
		TermsOfServiceComponent
	],
	imports: [
		AppRoutingModule,
		BrowserModule,
		HttpModule,
		FormsModule,
		ReactiveFormsModule,
		FlashMessagesModule.forRoot()
	],
	providers: [
		AuthGuard,
		ConfigService,
		FootballService,
		Title,
		UtilsService,
		{
			provide: APP_INITIALIZER,
			useFactory: ConfigLoader,
			deps: [ConfigService],
			multi: true
		}
	],
	bootstrap: [AppComponent]
})
export class AppModule { }