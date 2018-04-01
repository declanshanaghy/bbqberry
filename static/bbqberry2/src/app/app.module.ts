import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatSliderModule, MatCardModule} from '@angular/material';
import {HttpClientModule} from '@angular/common/http'

import {AppComponent} from './app.component';
import {TemperatureIndicatorComponent} from './temperature-indicator/temperature-indicator.component';
import {BbqberryService} from './bbqberry.service';
import {MessagesComponent} from './messages/messages.component';
import {MessageService} from './message.service';
import { LightstripComponent } from './lightstrip/lightstrip.component';

@NgModule({
  declarations: [
    AppComponent,
    TemperatureIndicatorComponent,
    MessagesComponent,
    LightstripComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatSliderModule, MatCardModule,
  ],
  providers: [
    BbqberryService,
    MessageService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
