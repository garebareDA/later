class IsLogin {
  is: Boolean;
  displayName:string|null;

  constructor(is:Boolean, displayName:string|null) {
    this.is = is;
    this.displayName = displayName;
  }
}

export default IsLogin;