import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LightstripComponent } from './lightstrip.component';

describe('LightstripComponent', () => {
  let component: LightstripComponent;
  let fixture: ComponentFixture<LightstripComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LightstripComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LightstripComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
