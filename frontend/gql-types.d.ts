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
  Cursor: any;
};

export type Tag = {
  __typename?: 'Tag';
  name: Scalars['String'];
  slug: Scalars['String'];
};

export enum MatchingPattern {
  PartialMatch = 'PARTIAL_MATCH',
  ExactMatch = 'EXACT_MATCH'
}

export type EdgeOrder = {
  key: OrderKey;
  direction: OrderDirection;
};

export type PageCondition = {
  backward?: Maybe<BackwardPagination>;
  forward?: Maybe<ForwardPagination>;
  nowPageNo: Scalars['Int'];
  initialLimit?: Maybe<Scalars['Int']>;
};

export type Node = {
  id: Scalars['ID'];
};

export enum OrderDirection {
  Asc = 'ASC',
  Desc = 'DESC'
}

export type PostEdge = Edge & {
  __typename?: 'PostEdge';
  node?: Maybe<Post>;
  cursor: Scalars['Cursor'];
};

export type Post = Node & {
  __typename?: 'Post';
  id: Scalars['ID'];
  post_date: Scalars['DateTime'];
  post_content: Scalars['String'];
  post_title: Scalars['String'];
  post_excerpt: Scalars['String'];
  post_modified: Scalars['DateTime'];
  category?: Maybe<Category>;
  tags?: Maybe<Array<Maybe<Tag>>>;
};

export type Connection = {
  pageInfo: PageInfo;
  edges: Array<Edge>;
  totalCount: Scalars['Int'];
};

export type Edge = {
  node?: Maybe<Node>;
  cursor: Scalars['Cursor'];
};

export enum PostOrderKey {
  PostDate = 'POST_DATE'
}


export type OrderKey = {
  postOrderKey?: Maybe<PostOrderKey>;
};

export type TextFilterCondition = {
  filterWord: Scalars['String'];
  matchingPattern?: Maybe<MatchingPattern>;
};

export type ForwardPagination = {
  first: Scalars['Int'];
  after?: Maybe<Scalars['Cursor']>;
};

export type PostConnection = Connection & {
  __typename?: 'PostConnection';
  pageInfo: PageInfo;
  edges: Array<PostEdge>;
  totalCount: Scalars['Int'];
};


export type Archive = {
  __typename?: 'Archive';
  year: Scalars['Int'];
  month: Scalars['Int'];
  posts: Scalars['Int'];
};

export type Category = {
  __typename?: 'Category';
  name: Scalars['String'];
  slug: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  posts: Array<Post>;
  postConnection?: Maybe<PostConnection>;
  archives?: Maybe<Array<Archive>>;
};


export type QueryPostConnectionArgs = {
  filterWord?: Maybe<TextFilterCondition>;
  pageCondition?: Maybe<PageCondition>;
  edgeOrder?: Maybe<EdgeOrder>;
};

export type BackwardPagination = {
  last: Scalars['Int'];
  before?: Maybe<Scalars['Cursor']>;
};

export type PageInfo = {
  __typename?: 'PageInfo';
  hasNextPage: Scalars['Boolean'];
  hasPreviousPage: Scalars['Boolean'];
  startCursor: Scalars['Cursor'];
  endCursor: Scalars['Cursor'];
};

