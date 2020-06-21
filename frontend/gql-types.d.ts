export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  DateTime: any;
};

export type Post = {
  __typename?: 'Post';
  id: Scalars['ID'];
  post_date: Scalars['DateTime'];
  post_content: Scalars['String'];
  post_title: Scalars['String'];
  post_excerpt: Scalars['String'];
  post_modified: Scalars['DateTime'];
};

export type Query = {
  __typename?: 'Query';
  posts: Array<Post>;
};


