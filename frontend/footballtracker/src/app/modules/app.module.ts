import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { AppRoutingModule } from '@app/app-routing.module';
import { HttpModule } from '@angular/http';

import { environment } from 'environments/environment';

// Pipes
import { CallbackPipe } from '@app/callback.pipe';

// Components
import { AppComponent } from './app.component';
import { FooterComponent } from '@components/footer/footer.component';
import { NavbarComponent } from '@components/navbar/navbar.component';

// Services
import { AuthGuard } from '@services/auth.guard';
import { ConfigService } from '@services/config.service';
import { FootballService } from '@services/football.service';

export function ConfigLoader(configService: ConfigService) {
	return () => configService.loadConfig(environment.httpConfig)
}

@NgModule({
	declarations: [
		AppComponent,
		CallbackPipe,
		FooterComponent,
		NavbarComponent
	],
	imports: [
		AppRoutingModule,
		BrowserModule,
		HttpModule
	],
	providers: [
		AuthGuard,
		ConfigService,
		FootballService,
		Title,
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
