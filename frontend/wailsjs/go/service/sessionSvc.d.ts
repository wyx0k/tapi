// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';
import {context} from '../models';

export function Close():Promise<void>;

export function GetTabSession():Promise<model.Response>;

export function Init(arg1:context.Context):Promise<void>;

export function SetTabSessionPosition(arg1:model.Tab,arg2:string):Promise<model.Response>;

export function SetTabSessionTabs(arg1:Array<model.Tab>):Promise<model.Response>;
