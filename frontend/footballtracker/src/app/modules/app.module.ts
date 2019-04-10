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
import { AppComponent } from './app.component';
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { FooterComponent } from '@components/footer/footer.component';
import { IndexComponent } from '@components/index/index.component';
import { LoginComponent } from '@components/login/login.component';
import { NavbarComponent } from '@components/navbar/navbar.component';
import { SignupComponent } from '@components/signup/signup.component';

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
		AppComponent,
		CallbackPipe,
		FixturesComponent,
		FooterComponent,
		IndexComponent,
		LoginComponent,
		NavbarComponent,
		SignupComponent
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