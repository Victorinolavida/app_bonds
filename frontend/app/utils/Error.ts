export class ErrorWithRequest extends Error {
  constructor(
    message: string,
    public uri: string
  ) {
    super(message);
    this.name = 'ErrorWithRequest';
    this.uri = uri;
  }
}
