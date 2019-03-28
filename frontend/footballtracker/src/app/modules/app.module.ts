import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from '@app/app-routing.module';

import { environment } from 'environments/environment';

// Pipes
import { CallbackPipe } from '@app/callback.pipe';

// Components
import { AppComponent } from './app.component';

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
		CallbackPipe
	],
	imports: [
		AppRoutingModule,
		BrowserModule
	],
	providers: [],
	bootstrap: [AppComponent]
})
export class AppModule { }
