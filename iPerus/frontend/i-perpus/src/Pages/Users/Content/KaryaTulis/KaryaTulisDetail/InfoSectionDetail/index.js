import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Button } from "../../ButtonElement";
import Novel from "../../../../../../novel.jpg";

import {
  Column2,
  Img,
  ImgWrap,
  InfoContainer,
  InfoWrapper,
  InfoRow,
  Column1,
  TextWrapper,
  TopLine,
  Heading,
  Subtitle,
  BtnWrap,
  Embed,
  ReadBook,
  Deskription,
} from "./InfoElements";

const InfoSectionDetail = ({
  lightBg,
  imgStart,
  topLine,
  lightText,
  headLine,
  darkText,
  description,
  buttonLabel,
  img,
  alt,
  id,
  primary,
  dark,
  dark2,
}) => {
  const params = useParams();
  const [article, setArticle] = useState({});
  const [loading, setLoading] = useState(true);
  const [notFound, setNotFound] = useState(false);

  useEffect(() => {
    const getArticles = async () => {
      const request = await fetch(`http://localhost:2213/api/bookList`);

      if (!request.ok) {
        return setNotFound(true);
      }

      const response = await request.json();
      console.log(response["AllBooks"][0]);

      setArticle(response["AllBooks"][0]);
      setLoading(false);
    };
    getArticles();
  }, [params]);

  if (notFound) {
    return <h1>Article Tidak Ditemukan :( </h1>;
  }

  return (
    <>
      <InfoContainer lightBg={lightBg} id={id}>
        <InfoWrapper>
          <InfoRow imgStart={imgStart}>
            <Column1>
              <TextWrapper>
                <TopLine>{article.Kategori}</TopLine>
                <Heading lightText={lightText}>{article.BookTitle}</Heading>
                <Subtitle darkText={darkText}>
                  {article.Writer}, {article.Tahun}
                </Subtitle>
                <Deskription darkText={darkText}>
                  Eu est duis duis ullamco nostrud velit ut Lorem fugiat quis
                  nisi. Ex est adipisicing laborum elit id cillum. Esse nostrud
                  in cupidatat nostrud et sint aliqua. Velit excepteur ullamco
                  nisi aliqua do exercitation fugiat non. Enim do nisi voluptate
                  anim culpa. Et sint elit dolor dolore aliquip nisi ea in id
                  sint adipisicing exercitation. Excepteur duis enim mollit eu
                  ea id non Lorem incididunt consectetur non cupidatat do sint.
                </Deskription>
                <ReadBook href={article.Link} target="_blank">
                  Baca Buku
                </ReadBook>
              </TextWrapper>
            </Column1>
            <Column2>
              <ImgWrap>
                <Img src={Novel} alt={alt} />
              </ImgWrap>
            </Column2>
          </InfoRow>
        </InfoWrapper>
      </InfoContainer>
    </>
  );
};

export default InfoSectionDetail;
