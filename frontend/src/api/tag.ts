import { fetchURL, removePrefix } from "./utils";

export async function getTag(url: string) {

  console.log("frontend/src/api/tag.ts/getTag()", url);

  url = removePrefix(url);

  const res = await fetchURL(`/api/tag${url}`, {});

  
  const data = (await res.json()) as Resource;
  console.log("after fetchURL", data);

  return res;
}





