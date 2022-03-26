import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WebsitesCheckerComponent } from './websites-checker.component';

describe('WebsitesCheckerComponent', () => {
  let component: WebsitesCheckerComponent;
  let fixture: ComponentFixture<WebsitesCheckerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WebsitesCheckerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WebsitesCheckerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
