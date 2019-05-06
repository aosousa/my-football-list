import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { AppRoutingModule } from '@app/app-routing.module';
import { HttpModule } from '@angular/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { PlatformModule } from '@angular/cdk/platform';
import { DataTablesModule } from 'angular-datatables';

import { environment } from 'environments/environment';

// 3rd party
import { FlashMessagesModule } from 'angular2-flash-messages';

// Pipes
import { CallbackPipe } from '@app/callback.pipe';

// Components
import { AboutComponent } from '@components/about/about.component';
import { AppComponent } from './app.component';
import { ChangePasswordComponent } from '@components/change-password/change-password.component';
import { ContactComponent } from '@components/contact/contact.component';
import { EditProfileComponent } from '@components/edit-profile/edit-profile.component';
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { FooterComponent } from '@components/footer/footer.component';
import { IndexComponent } from '@components/index/index.component';
import { LeagueComponent } from '@components/league/league.component';
import { LoginComponent } from '@components/login/login.component';
import { NavbarComponent } from '@components/navbar/navbar.component';
import { PrivacyPolicyComponent } from '@components/privacy-policy/privacy-policy.component';
import { ProfileComponent } from '@components/profile/profile.component';
import { ResetPasswordComponent } from '@components/reset-password/step1/reset-password.component';
import { NewPasswordComponent } from '@components/reset-password/step2/new-password.component';
import { SignupComponent } from '@components/signup/signup.component';
import { TeamComponent } from '@components/team/team.component';
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
		ChangePasswordComponent,
		ContactComponent,
		EditProfileComponent,
		FixturesComponent,
		FooterComponent,
		IndexComponent,
		LeagueComponent,
		LoginComponent,
		NavbarComponent,
		PrivacyPolicyComponent,
		ProfileComponent,
		ResetPasswordComponent,
		NewPasswordComponent,
		SignupComponent,
		TeamComponent,
		TermsOfServiceComponent
	],
	imports: [
		AppRoutingModule,
		BrowserModule,
		DataTablesModule,
		HttpModule,
		FormsModule,
		ReactiveFormsModule,
		FlashMessagesModule.forRoot(),
		NgbModule,
		ScrollingModule,
		PlatformModule
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