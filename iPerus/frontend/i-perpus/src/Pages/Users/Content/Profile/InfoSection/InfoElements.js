import styled from "styled-components";

export const InfoContainer = styled.div`
  color: black;
  background: white;
  margin-top: 50px;
  margin-bottom: 100px;

  @media screen and (max-width: 768px) {
    padding: 100px 0;
  }
`;

export const InfoWrapper = styled.div`
  display: grid;
  z-index: 1;
  width: 100%;
  margin-right: auto;
  margin-left: auto;
  padding: 0 24px;
  justify-content: center;
`;
