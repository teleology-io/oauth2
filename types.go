package oauth2

type Scope string | []string;

func Test() Scope {
	return "hello"
}

func Test2() Scope {
	it := []string{"hello", "world"}
}

// export interface Token {
//   token: string,
//   user_id?: string,
//   client_id?: string,
//   scope?: Scope,
//   created_at: Date,
//   [key: string]: any;
// }

// export interface Code {
//   code: string,
//   user_id?: string,
//   client_id?: string,
//   scope?: Scope,
//   created_at: Date,
//   [key: string]: any;
// };

// export interface AccessToken extends Token {};

// export interface RefreshToken extends Token {};

// export type TokenRequest = {
//   user_id?: string,
//   client_id?: string,
//   scope?: Scope,
//   ttl: number
// };

// export type User = {
//   id: string,
//   username: string,
//   password?: string,
//   [key: string]: any;
// }

// export type Client = {
//   id: string,
//   secret?: string,
//   grants: string[],
//   redirect_uris: string[],
//   scopes: string[];
//   [key: string]: any;
// }

// export type Body = string | {
//   [key: string]: any;
// }

// export type Query = {
//   [key: string]: any;
// }

// export type Headers = {
//   [key: string]: any;
// }

// export interface Request {
//   method: 'get' | 'post' | 'put' | 'delete' | 'options' | 'patch'
//   headers: Headers, 
//   body?: Body,
//   query?: Query,
// }

// export interface DecisionRequest extends Request {
//   client_id: string,
//   scope: Scope,
// };

// export interface Response {
//   code: number,
//   headers: Headers, 
//   body?: Body,
// }

// export interface IntrospectionResponse {
//   active: boolean,
//   client_id?: string,
//   username?: string,
//   token_type?: string,
//   exp?: number,
//   iat?: number,
//   nbf?: number,
//   sub?:string,
//   aud?: string,
//   iss?: string,
//   jti?: string,
// }

// export interface Options {
//   codeTtl: number,
//   accessTokenTtl: number,
//   refreshTokenTtl: number,

//   // Decision
//   createDecisionPage(req: DecisionRequest): Promise<string>;
//   createCode(req: TokenRequest): Promise<Code>;
//   createAccessToken(req: TokenRequest): Promise<AccessToken>;
//   createRefreshToken(req: TokenRequest): Promise<AccessToken>;

//   // Token
//   getTokenTtl(token: Token): number;
//   getCode(code: string): Promise<Code>;
//   getAccessTokenWithIds(user_id: string, client_id: string): Promise<AccessToken>;
//   getAccessToken(token: string): Promise<AccessToken>;
//   getRefreshToken(token: string): Promise<RefreshToken>;
//   introspect(token: Token): IntrospectionResponse;

//   // Client
//   getClient(id: string): Promise<Client>;
//   validGrantType(client: Client, grant_type: string): boolean;
//   validSecret(client: Client, client_secret: string): boolean;
//   validScope(client: Client, scope: Scope): boolean;
//   validRedirectUri(client: Client, redirect_uri: string): boolean;

//   // User
//   getUser(id: string): Promise<User>;
//   getUserByName(username: string): Promise<User>;
//   validPassword(user: User, password: string): boolean;

// }

// export interface Oauth2 {

//   // Endpoints
//   token(req: Request): Promise<Response>;
//   authorize(req: Request): Promise<Response>;
//   introspection(req: Request): Promise<Response>;
  
// }
